import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { combineLatest, Subject } from 'rxjs';
import { filter, pluck, takeUntil } from 'rxjs/operators';
import { identity, isNil } from 'lodash/fp';

import { LayoutFacadeService, Sidebar } from 'app/entities/layout/layout.facade';
import { routeParams, routeURL } from 'app/route.selectors';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { GetNodeCredential, UpdateNodeCredential } from 'app/entities/node-credentials/node-credential.actions';
import { credentialFromRoute, getStatus } from 'app/entities/node-credentials/node-credential-details.selectors';
import { updateStatus } from 'app/entities/node-credentials/node-credential.selectors';
import { NodeCredential, SaveNodeCredential } from 'app/entities/node-credentials/node-credential.model';
import { pending, EntityStatus, allLoaded } from 'app/entities/entities';

export type NodeCredentialTabName = 'details' | 'reset';

@Component({
  selector: 'app-node-credential-details-screen',
  templateUrl: './node-credential-details.component.html',
  styleUrls: ['./node-credential-details.component.scss']
})

export class NodeCredentialDetailsScreenComponent implements OnInit {
  private isDestroyed = new Subject<boolean>();
  public nodeCredential: NodeCredential;
  public tabValue: NodeCredentialTabName = 'details';
  public selected = 'pass';
  public types: string[] = ['Password', 'RSA'];
  public passwordSelected = 'Password';
  public updateForm: FormGroup;
  public resetForm: FormGroup;
  public sshForms: FormGroup;
  public winrmForms: FormGroup;
  public sudoForms: FormGroup;
  public id: string;
  public url: string;
  public saveSuccessful = false;
  public saveInProgress = false;
  public resetSuccessful = false;
  public resetInProgress = false;
  public isLoading = true;
  nodeDetailsLoading = true;

  constructor(
    private store: Store<NgrxStateAtom>,
    private saveCred: SaveNodeCredential,
    private fb: FormBuilder,
    private router: Router,
    private layoutFacade: LayoutFacadeService
  ) {
    this.sshForms = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
      key: ['', Validators.required]
    });
    this.winrmForms = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    this.sudoForms = this.fb.group({
      password: ['', Validators.required],
      options: ['', Validators.required]
    });

    this.updateForm = this.fb.group({
      id: [''],
      name: ['', Validators.required],
      type: ['', Validators.required],
      ssh: this.sshForms,
      winrm: this.winrmForms,
      sudo: this.sudoForms
    });

    this.resetForm = this.fb.group({
      id: [''],
      name: ['', Validators.required],
      type: ['', Validators.required],
      ssh: this.sshForms,
      winrm: this.winrmForms,
      sudo: this.sudoForms
    });
    this.getNodeCred();
  }

  ngOnInit() {
    this.layoutFacade.showSidebar(Sidebar.Settings);
    this.store.select(routeURL).pipe(takeUntil(this.isDestroyed))
      .subscribe((url: string) => {
        this.url = url;
        const [, fragment] = url.split('#');
        this.tabValue = (fragment === 'reset') ? 'reset' : 'details';
      });

      this.store.select(updateStatus).pipe(
        takeUntil(this.isDestroyed),
        filter(state => this.saveInProgress || this.resetInProgress && !pending(state)))
        .subscribe((state) => {
          if (this.tabValue === 'details') {
          this.saveInProgress = false;
          this.saveSuccessful = (state === EntityStatus.loadingSuccess);
          this.updateForm.markAsPristine();
          } else {
          this.resetInProgress = false;
          this.resetSuccessful = (state === EntityStatus.loadingSuccess);
          this.resetForm.markAsPristine();
          }
          this.getNodeCred();
        });
  }


  onSelectedTab(event: { target: { value: NodeCredentialTabName } }) {
    this.tabValue = event.target.value;
    this.router.navigate([this.url.split('#')[0]], { fragment: event.target.value });
  }

  selectChangeHandlers(id: string): void {
    this.passwordSelected = id;
  }

  saveNodeCredential(data): void {
    let credBody;
    if (this.tabValue === 'details') {
      this.saveSuccessful = false;
      this.saveInProgress = true;
      credBody = {
        id: data['id'],
        name: data['name'],
        type: this.nodeCredential.type,
        data: this.nodeCredential.data,
        tags: this.nodeCredential.tags
      };
    } else {
      this.resetSuccessful = false;
      this.resetInProgress = true;
      credBody = this.saveCred.getNodeCredentialCreate(data);
    }

    this.store.dispatch(new UpdateNodeCredential(credBody));
    this.resetForm.reset();
  }

  getNodeCred() {
    combineLatest([
      this.store.select(routeParams).pipe(pluck('id'), filter(identity))
    ]).pipe(
      takeUntil(this.isDestroyed)
    ).subscribe(([id]: string[]) => {
      this.id = id;
      this.store.dispatch(new GetNodeCredential({
        id: id
      }));
    });

    combineLatest([
      this.store.select(getStatus),
      this.store.select(credentialFromRoute)]).pipe(
      filter(([status, nodeCredential]) =>
      status === EntityStatus.loadingSuccess && !isNil(nodeCredential)),
      takeUntil(this.isDestroyed))
      .subscribe(([status, nodeCredential]) => {
        this.isLoading =
        !allLoaded([status]);
      this.nodeCredential = nodeCredential;
      const data = this.nodeCredential.data;

      setTimeout(() => (
      this.updateForm.controls.name.setValue(this.nodeCredential.name),
      this.updateForm.controls.type.setValue(this.nodeCredential.type),
      this.updateForm.controls.id.setValue(this.nodeCredential.id)
      ),
      200);
      if (this.nodeCredential.type === 'sudo') {
        for (let i = 0; i < data.length; i++) {
          if (data[i].key === 'options') {
            this.sudoForms.controls.options.setValue(data[i].value);
          }
        }
      }
      this.resetForm.controls.name.setValue(this.nodeCredential.name);
      this.resetForm.controls.type.setValue(this.nodeCredential.type);
      this.resetForm.controls.id.setValue(this.nodeCredential.id);
      this.nodeDetailsLoading = false;
    });

  }
}
