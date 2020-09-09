import { Component, OnInit, OnDestroy, EventEmitter } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { get, toUpper, pipe, set, pick } from 'lodash/fp';
import { Observable, Subject } from 'rxjs';
import { map } from 'rxjs/operators';
import * as moment from 'moment/moment';
import { MatOptionSelectionChange } from '@angular/material/core';

import { NgrxStateAtom } from 'app/ngrx.reducers';
import { LayoutFacadeService, Sidebar } from 'app/entities/layout/layout.facade';
import { DateTime } from 'app/helpers/datetime/datetime';
import { NodeCredentialsSearch, DeleteNodeCredential } from 'app/entities/node-credentials/node-credential.actions';
import {  loading } from 'app/entities/entities';
import {
  getAllStatus
 } from 'app/entities/node-credentials/node-credential.selectors';
import { NodeCredential, NodeCredentialTypes } from 'app/entities/node-credentials/node-credential.model';
import {
  allCredentials
} from 'app/entities/node-credentials/node-credential.selectors';

import { NodeCredentialOrder } from './node-credential-list.reducer';
import { nodeCredentialListState } from './node-credential-list.selectors';
import { SortNodeCredentialList } from './node-credential-list.actions';
@Component({
  selector: 'app-node-credential-list',
  templateUrl: './node-credential-list.component.html',
  styleUrls: ['./node-credential-list.component.scss']
})
export class  NodeCredentialListComponent implements OnInit, OnDestroy {
  private isDestroyed = new Subject<boolean>();
  public loading$: Observable<boolean>;
  public instanceNodeCredentials$: Observable<NodeCredential[]>;
  public conflictErrorEvent = new EventEmitter<boolean>();
  public openUserModal = new EventEmitter<boolean>();
  public list: NodeCredential[] = [];
  public orderBy: NodeCredentialOrder;
  public nodeCredentialToDelete: NodeCredential;
  public createModalVisible = false;
  public creatingNodeCredential = false;
  public deleteModalVisible = false;
  public nodesListLoading = true;
  public messageModalVisible = false;
  public nodesListSort = 'last_contact';
  public nodesListSortOrder = 'DESC';
  public nodesListFilters = [{key: 'type', values: ['ssh', 'winrm', 'sudo']}];
  public sortBy: string;
  public params: object;

  constructor(
    private store: Store<NgrxStateAtom>,
    private layoutFacade: LayoutFacadeService
  ) {
    this.loading$ = store.select(getAllStatus).pipe(map(loading));
  }

  ngOnInit(): void {
    this.layoutFacade.showSidebar(Sidebar.Settings);
    this.getNodeList();

  }

  ngOnDestroy() {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }

  public openCreateModal(): void {
    this.openUserModal.emit(true);
  }

  orderFor(column) {
    return column === this.sortBy ? this.orderBy : 'none';
  }

  handleSortToggle({detail: sortParams}) {

    this.store.dispatch(new SortNodeCredentialList(sortParams));
    this.getNodeList();
  }

  formatKeyType(credentialType: NodeCredentialTypes): string {
    switch (credentialType) {
      case NodeCredentialTypes.SSH:
        return 'SSH';
      case NodeCredentialTypes.WinRM:
        return 'WinRM';
      case NodeCredentialTypes.Sudo:
        return 'Sudo';
      default:
        return 'SSH';
    }
  }

  public startNodeCredentialDelete( $event: MatOptionSelectionChange, p: NodeCredential): void {
    if ($event.isUserInput) {
      this.nodeCredentialToDelete = p;
      this.deleteModalVisible = true;
    }
  }

  getNodeList() {
    this.instanceNodeCredentials$ = this.store.pipe(select(allCredentials));
    this.store.select(nodeCredentialListState).subscribe((state) => {
      this.params = pick(['filters', 'page', 'per_page'], state);

      this.sortBy = get('sort', state);
      this.orderBy = get('order', state);

      if (this.orderBy !== 'none') {
        this.params = pipe(
          set('order', toUpper(this.orderBy)),
          set('sort', this.sortBy)
        )(this.params) as object;
      }
    });
    this.store.dispatch(new NodeCredentialsSearch(this.params));
    this.nodesListLoading = false;
  }

  public deleteNodeCredential(): void {
    this.closeDeleteModal();
    this.store.dispatch(new DeleteNodeCredential(this.nodeCredentialToDelete));
  }

  public closeDeleteModal(): void {
    this.deleteModalVisible = false;
  }


  public closeMessageModal(): void {
    this.messageModalVisible = false;
  }


  public getLastModified(lastModifiedDate): string {
    return moment.utc(lastModifiedDate).format(DateTime.RFC2822);
  }

  public inUseMessage(): string {
    return (this.nodeCredentialToDelete && this.nodeCredentialToDelete.data.length > 0) ?
      'and will disrupt access for all attached members' : '';
  }
}
