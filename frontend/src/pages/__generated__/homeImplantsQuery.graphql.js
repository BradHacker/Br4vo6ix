/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type homeImplantsQueryVariables = {||};
export type homeImplantsQueryResponse = {|
  +implants: $ReadOnlyArray<{|
    +uuid: string,
    +machine_id: string,
    +last_seen_at: ?any,
    +tasks: $ReadOnlyArray<{|
      +uuid: string,
      +type: TaskType,
      +payload: string,
      +stdout: ?string,
      +stderr: ?string,
      +created_at: any,
    |}>,
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
    machine_id
    last_seen_at
    tasks {
      uuid
      type
      payload
      stdout
      stderr
      created_at
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
        "kind": "ScalarField",
        "name": "last_seen_at",
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
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "created_at",
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
    "name": "homeImplantsQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "homeImplantsQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "28265490c2a51be795aba052d95f1f6a",
    "id": null,
    "metadata": {},
    "name": "homeImplantsQuery",
    "operationKind": "query",
    "text": "query homeImplantsQuery {\n  implants {\n    uuid\n    machine_id\n    last_seen_at\n    tasks {\n      uuid\n      type\n      payload\n      stdout\n      stderr\n      created_at\n    }\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = 'b0a36928bf206e7fd98f40323ba33a19';

module.exports = node;
