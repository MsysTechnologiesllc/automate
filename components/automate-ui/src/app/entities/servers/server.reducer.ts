import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { HttpErrorResponse } from '@angular/common/http';
import { pipe, set, unset } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import { ServerActionTypes, ServerActions } from './server.actions';
import { Server } from './server.model';

export interface ServerEntityState extends EntityState<Server> {
  status: EntityStatus;
  saveStatus: EntityStatus;
  saveError: HttpErrorResponse;
  getStatus: EntityStatus;
  deleteStatus: EntityStatus;
}

// reusable names of the above properties:
// TODO (tc): Rename this to GET_ALL_STATUS
const STATUS = 'status';
const SAVE_STATUS = 'saveStatus';
const SAVE_ERROR = 'saveError';
const GET_STATUS = 'getStatus';

export const serverEntityAdapter: EntityAdapter<Server> = createEntityAdapter<Server>();

export const ServerEntityInitialState: ServerEntityState =
  serverEntityAdapter.getInitialState({
    status: EntityStatus.notLoaded,
    saveStatus: EntityStatus.notLoaded,
    saveError: null,
    getStatus: EntityStatus.notLoaded,
    deleteStatus: EntityStatus.notLoaded
  });

export function serverEntityReducer(
  state: ServerEntityState = ServerEntityInitialState,
  action: ServerActions): ServerEntityState {

  switch (action.type) {

    case ServerActionTypes.GET_ALL: {
      return set(
        STATUS,
        EntityStatus.loading,
        serverEntityAdapter.removeAll(state)
      ) as ServerEntityState;
    }

    case ServerActionTypes.GET_ALL_SUCCESS: {
      return set(
        STATUS,
        EntityStatus.loadingSuccess,
        serverEntityAdapter.addAll(action.payload.servers, state)
      ) as ServerEntityState;
    }

    case ServerActionTypes.GET_ALL_FAILURE: {
      return set(STATUS, EntityStatus.loadingFailure, state) as ServerEntityState;
    }

    case ServerActionTypes.GET: {
      return set(
        GET_STATUS,
        EntityStatus.loading,
        serverEntityAdapter.removeAll(state)
      ) as ServerEntityState;
    }

    case ServerActionTypes.GET_SUCCESS: {
      return set(
        GET_STATUS,
        EntityStatus.loadingSuccess,
        serverEntityAdapter.addOne(action.payload.server, state)
      ) as ServerEntityState;
    }

    case ServerActionTypes.GET_FAILURE: {
      return set(GET_STATUS, EntityStatus.loadingFailure, state) as ServerEntityState;
    }

    case ServerActionTypes.CREATE: {
      return set(SAVE_STATUS, EntityStatus.loading, state) as ServerEntityState;
    }

    case ServerActionTypes.CREATE_SUCCESS: {
      return pipe(
        unset(SAVE_ERROR),
        set(SAVE_STATUS, EntityStatus.loadingSuccess)
      )(serverEntityAdapter.addOne(action.payload.server, state)) as ServerEntityState;
    }

    case ServerActionTypes.CREATE_FAILURE: {
      return pipe(
        set(SAVE_ERROR, action.payload),
        set(SAVE_STATUS, EntityStatus.loadingFailure)
      )(state) as ServerEntityState;
    }

    case ServerActionTypes.DELETE: {
      return set('deleteStatus', EntityStatus.loading, state);
    }

    case ServerActionTypes.DELETE_SUCCESS: {
      return set('deleteStatus', EntityStatus.loadingSuccess,
        serverEntityAdapter.removeOne(action.payload.id, state));
    }

    case ServerActionTypes.DELETE_FAILURE: {
      return set('deleteStatus', EntityStatus.loadingFailure, state);
    }

  }

  return state;
}
