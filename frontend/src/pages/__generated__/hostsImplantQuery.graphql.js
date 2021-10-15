/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type hostsImplantQueryVariables = {|
  implantId: string
|};
export type hostsImplantQueryResponse = {|
  +implant: {|
    +uuid: string,
    +hostname: string,
    +ip: string,
    +machine_id: string,
    +last_seen_at: ?any,
  |}
|};
export type hostsImplantQuery = {|
  variables: hostsImplantQueryVariables,
  response: hostsImplantQueryResponse,
|};
*/


/*
query hostsImplantQuery(
  $implantId: String!
) {
  implant(implantUuid: $implantId) {
    uuid
    hostname
    ip
    machine_id
    last_seen_at
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "implantId"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "implantUuid",
        "variableName": "implantId"
      }
    ],
    "concreteType": "Implant",
    "kind": "LinkedField",
    "name": "implant",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "uuid",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "hostname",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "ip",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "machine_id",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "last_seen_at",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "hostsImplantQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "hostsImplantQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "6c6e139250df40aa87b28aca74b20dbc",
    "id": null,
    "metadata": {},
    "name": "hostsImplantQuery",
    "operationKind": "query",
    "text": "query hostsImplantQuery(\n  $implantId: String!\n) {\n  implant(implantUuid: $implantId) {\n    uuid\n    hostname\n    ip\n    machine_id\n    last_seen_at\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '9f36c0780dc146fdb641089a8dafc637';

module.exports = node;
