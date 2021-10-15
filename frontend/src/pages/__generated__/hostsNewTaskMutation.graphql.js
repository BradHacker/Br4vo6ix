/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type hostsNewTaskMutationVariables = {|
  implantId: string,
  type: TaskType,
  payload: string,
|};
export type hostsNewTaskMutationResponse = {|
  +scheduleTask: {|
    +uuid: string
  |}
|};
export type hostsNewTaskMutation = {|
  variables: hostsNewTaskMutationVariables,
  response: hostsNewTaskMutationResponse,
|};
*/


/*
mutation hostsNewTaskMutation(
  $implantId: String!
  $type: TaskType!
  $payload: String!
) {
  scheduleTask(input: {implantUuid: $implantId, type: $type, payload: $payload}) {
    uuid
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "implantId"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "payload"
},
v2 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "type"
},
v3 = [
  {
    "alias": null,
    "args": [
      {
        "fields": [
          {
            "kind": "Variable",
            "name": "implantUuid",
            "variableName": "implantId"
          },
          {
            "kind": "Variable",
            "name": "payload",
            "variableName": "payload"
          },
          {
            "kind": "Variable",
            "name": "type",
            "variableName": "type"
          }
        ],
        "kind": "ObjectValue",
        "name": "input"
      }
    ],
    "concreteType": "Task",
    "kind": "LinkedField",
    "name": "scheduleTask",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "uuid",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*: any*/),
      (v1/*: any*/),
      (v2/*: any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "hostsNewTaskMutation",
    "selections": (v3/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [
      (v0/*: any*/),
      (v2/*: any*/),
      (v1/*: any*/)
    ],
    "kind": "Operation",
    "name": "hostsNewTaskMutation",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "10f4f69b12648f2016c7eb4f6cf6954b",
    "id": null,
    "metadata": {},
    "name": "hostsNewTaskMutation",
    "operationKind": "mutation",
    "text": "mutation hostsNewTaskMutation(\n  $implantId: String!\n  $type: TaskType!\n  $payload: String!\n) {\n  scheduleTask(input: {implantUuid: $implantId, type: $type, payload: $payload}) {\n    uuid\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = 'c3542f15edc742a2ed06a042a6f8bbcd';

module.exports = node;
