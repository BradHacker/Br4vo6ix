/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type homeNewTaskMutationVariables = {|
  implantId: string,
  type: TaskType,
  payload: string,
|};
export type homeNewTaskMutationResponse = {|
  +scheduleTask: {|
    +uuid: string
  |}
|};
export type homeNewTaskMutation = {|
  variables: homeNewTaskMutationVariables,
  response: homeNewTaskMutationResponse,
|};
*/


/*
mutation homeNewTaskMutation(
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
    "name": "homeNewTaskMutation",
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
    "name": "homeNewTaskMutation",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "10610b3ed2ec03fc3e3d0a317c872848",
    "id": null,
    "metadata": {},
    "name": "homeNewTaskMutation",
    "operationKind": "mutation",
    "text": "mutation homeNewTaskMutation(\n  $implantId: String!\n  $type: TaskType!\n  $payload: String!\n) {\n  scheduleTask(input: {implantUuid: $implantId, type: $type, payload: $payload}) {\n    uuid\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '69dd15864f51d264da25502fbbee6ee4';

module.exports = node;
