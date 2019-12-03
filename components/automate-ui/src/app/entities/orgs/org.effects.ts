import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of as observableOf } from 'rxjs';
import { catchError, mergeMap, map, filter } from 'rxjs/operators';

import { HttpStatus } from 'app/types/types';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';

import {
  GetOrgsForServer,
  GetOrgsSuccess,
  OrgsSuccessPayload,
  GetOrgsFailure,
  GetOrg,
  GetOrgSuccess,
  GetOrgFailure,
  CreateOrg,
  CreateOrgSuccess,
  CreateOrgFailure,
  DeleteOrg,
  DeleteOrgSuccess,
  DeleteOrgFailure,
  UpdateOrg,
  UpdateOrgFailure,
  UpdateOrgSuccess,
  OrgSuccessPayload,
  OrgActionTypes
} from './org.actions';

import {
  OrgRequests
} from './org.requests';

@Injectable()
export class OrgEffects {
  constructor(
    private actions$: Actions,
    private requests: OrgRequests
  ) { }

  @Effect()
  getOrgsForProject$ = this.actions$.pipe(
      ofType(OrgActionTypes.GET_ALL),
      mergeMap(({ payload: { server_id } }: GetOrgsForServer) =>
        this.requests.getOrgsForServer(server_id).pipe(
          map((resp: OrgsSuccessPayload) => new GetOrgsSuccess(resp)),
          catchError((error: HttpErrorResponse) => observableOf(new GetOrgsFailure(error))))));

  @Effect()
  getOrgsFailure$ = this.actions$.pipe(
      ofType(OrgActionTypes.GET_ALL_FAILURE),
      map(({ payload }: GetOrgsFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get orgs: ${msg || payload.error}`
        });
      }));

  @Effect()
  getOrg$ = this.actions$.pipe(
      ofType(OrgActionTypes.GET),
      mergeMap(({ payload: { server_id, id } }: GetOrg) =>
        this.requests.getOrg(server_id, id).pipe(
          map((resp: OrgSuccessPayload) => new GetOrgSuccess(resp)),
          catchError((error: HttpErrorResponse) => observableOf(new GetOrgFailure(error))))));

 @Effect()
  getOrgFailure$ = this.actions$.pipe(
      ofType(OrgActionTypes.GET_FAILURE),
      map(({ payload }: GetOrgFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get org: ${msg || payload.error}`
        });
      }));

  @Effect()
  createOrg$ = this.actions$.pipe(
      ofType(OrgActionTypes.CREATE),
      mergeMap(({ payload }: CreateOrg) =>
      this.requests.createOrg( payload ).pipe(
        map((resp: OrgSuccessPayload) => new CreateOrgSuccess(resp)),
        catchError((error: HttpErrorResponse) => observableOf(new GetOrgFailure(error))))));

  @Effect()
  createOrgSuccess$ = this.actions$.pipe(
      ofType(OrgActionTypes.CREATE_SUCCESS),
      map(({ payload: { org } }: CreateOrgSuccess) => new CreateNotification({
      type: Type.info,
      message: `Created org ${org.id}`
    })));

  @Effect()
  createOrgFailure$ = this.actions$.pipe(
    ofType(OrgActionTypes.CREATE_FAILURE),
    filter(({ payload }: CreateOrgFailure) => payload.status !== HttpStatus.CONFLICT),
    map(({ payload }: CreateOrgFailure) => new CreateNotification({
        type: Type.error,
        message: `Could not create org: ${payload.error.error || payload}`
      })));

  @Effect()
  deleteOrg$ = this.actions$.pipe(
      ofType(OrgActionTypes.DELETE),
      mergeMap(({ payload: { server_id, id } }: DeleteOrg) =>
        this.requests.deleteOrg(server_id, id).pipe(
          map(() => new DeleteOrgSuccess({id})),
          catchError((error: HttpErrorResponse) =>
            observableOf(new DeleteOrgFailure(error))))));

  @Effect()
  deleteOrgSuccess$ = this.actions$.pipe(
      ofType(OrgActionTypes.DELETE_SUCCESS),
      map(({ payload: { id } }: DeleteOrgSuccess) => {
        return new CreateNotification({
          type: Type.info,
          message: `Deleted org ${id}.`
        });
      }));

  @Effect()
  deleteOrgFailure$ = this.actions$.pipe(
      ofType(OrgActionTypes.DELETE_FAILURE),
      map(({ payload }: DeleteOrgFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not delete org: ${msg || payload.error}`
        });
      }));

  @Effect()
  updateOrg$ = this.actions$.pipe(
      ofType(OrgActionTypes.UPDATE),
      mergeMap(({ payload: { org } }: UpdateOrg) =>
        this.requests.updateOrg(org).pipe(
          map((resp: OrgSuccessPayload) => new UpdateOrgSuccess(resp)),
          catchError((error: HttpErrorResponse) =>
            observableOf(new UpdateOrgFailure(error))))));

  @Effect()
  updateOrgFailure$ = this.actions$.pipe(
      ofType(OrgActionTypes.UPDATE_FAILURE),
      map(({ payload }: UpdateOrgFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not update org: ${msg || payload.error}`
        });
      }));
}

