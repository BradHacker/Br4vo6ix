/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type AppImplantsQueryVariables = {||};
export type AppImplantsQueryResponse = {|
  +implants: $ReadOnlyArray<{|
    +uuid: string,
    +machine_id: string,
    +tasks: $ReadOnlyArray<{|
      +uuid: string,
      +type: TaskType,
      +payload: string,
      +stdout: ?string,
      +stderr: ?string,
    |}>,
  |}>
|};
export type AppImplantsQuery = {|
  variables: AppImplantsQueryVariables,
  response: AppImplantsQueryResponse,
|};
*/


/*
query AppImplantsQuery {
  implants {
    uuid
    machine_id
    tasks {
      uuid
      type
      payload
      stdout
      stderr
    }
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "uuid",
  "storageKey": null
},
v1 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "Implant",
    "kind": "LinkedField",
    "name": "implants",
    "plural": true,
    "selections": [
      (v0/*: any*/),
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
        "concreteType": "Task",
        "kind": "LinkedField",
        "name": "tasks",
        "plural": true,
        "selections": [
          (v0/*: any*/),
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "type",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "payload",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "stdout",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "stderr",
            "storageKey": null
          }
        ],
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
    "name": "AppImplantsQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "AppImplantsQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "c771ca80f29c7e9dc5a428a2ea711aa1",
    "id": null,
    "metadata": {},
    "name": "AppImplantsQuery",
    "operationKind": "query",
    "text": "query AppImplantsQuery {\n  implants {\n    uuid\n    machine_id\n    tasks {\n      uuid\n      type\n      payload\n      stdout\n      stderr\n    }\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = 'e292c10dc42770a726bcf28188893680';

module.exports = node;
