import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';

import { Server } from './server.model';

export enum ServerActionTypes {
  GET_ALL                        = 'SERVER::GET_ALL',
  GET_ALL_SUCCESS                = 'SERVER::GET_ALL::SUCCESS',
  GET_ALL_FAILURE                = 'SERVER::GET_ALL::FAILURE',
  GET                            = 'SERVER::GET',
  GET_SUCCESS                    = 'SERVER::GET::SUCCESS',
  GET_FAILURE                    = 'SERVER::GET::FAILURE',
  CREATE                         = 'SERVER::CREATE',
  CREATE_SUCCESS                 = 'SERVER::CREATE::SUCCESS',
  CREATE_FAILURE                 = 'SERVER::CREATE::FAILURE'
}


export interface ServerSuccessPayload {
  server: Server;
}

export interface GetServersSuccessPayload {
  servers: Server[];
}

export class GetServers implements Action {
  readonly type = ServerActionTypes.GET_ALL;
}

export class GetServersSuccess implements Action {
  readonly type = ServerActionTypes.GET_ALL_SUCCESS;

  constructor(public payload: GetServersSuccessPayload) { }
}

export class GetServersFailure implements Action {
  readonly type = ServerActionTypes.GET_ALL_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class GetServer implements Action {
  readonly type = ServerActionTypes.GET;

  constructor(public payload: { id: string }) { }
}

export class GetServerSuccess implements Action {
  readonly type = ServerActionTypes.GET_SUCCESS;

  constructor(public payload: ServerSuccessPayload) { }
}

export class GetServerFailure implements Action {
  readonly type = ServerActionTypes.GET_FAILURE;

  constructor(public payload: HttpErrorResponse, public id: string) { }
}

export interface CreateServerPayload {
  name: string;
  description: string;
  fqdn: string;
  ip_address: string;
}

export class CreateServer implements Action {
  readonly type = ServerActionTypes.CREATE;
  constructor(public payload: CreateServerPayload) { }
}

export class CreateServerSuccess implements Action {
  readonly type = ServerActionTypes.CREATE_SUCCESS;
  constructor(public payload: ServerSuccessPayload) { }
}

export class CreateServerFailure implements Action {
  readonly type = ServerActionTypes.CREATE_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export type ServerActions =
  | GetServers
  | GetServersSuccess
  | GetServersFailure
  | GetServer
  | GetServerSuccess
  | GetServerFailure
  | CreateServer
  | CreateServerSuccess
  | CreateServerFailure;
