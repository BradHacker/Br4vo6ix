/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type homeImplantsQueryVariables = {||};
export type homeImplantsQueryResponse = {|
  +implants: $ReadOnlyArray<{|
    +uuid: string,
    +hostname: string,
    +ip: string,
    +machine_id: string,
    +last_seen_at: ?any,
  |}>
|};
export type homeImplantsQuery = {|
  variables: homeImplantsQueryVariables,
  response: homeImplantsQueryResponse,
|};
*/


/*
query homeImplantsQuery {
  implants {
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
    "alias": null,
    "args": null,
    "concreteType": "Implant",
    "kind": "LinkedField",
    "name": "implants",
    "plural": true,
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
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "homeImplantsQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "homeImplantsQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "ceea0c57945a139b5a86b3e41e5be3de",
    "id": null,
    "metadata": {},
    "name": "homeImplantsQuery",
    "operationKind": "query",
    "text": "query homeImplantsQuery {\n  implants {\n    uuid\n    hostname\n    ip\n    machine_id\n    last_seen_at\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = 'eca9f8c851d36eca8e9d59a4047875b9';

module.exports = node;
