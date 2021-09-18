/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type tasksGetTasksQueryVariables = {|
  implantId: string
|};
export type tasksGetTasksQueryResponse = {|
  +tasks: $ReadOnlyArray<{|
    +uuid: string,
    +type: TaskType,
    +payload: string,
    +stdout: ?string,
    +stderr: ?string,
  |}>
|};
export type tasksGetTasksQuery = {|
  variables: tasksGetTasksQueryVariables,
  response: tasksGetTasksQueryResponse,
|};
*/


/*
query tasksGetTasksQuery(
  $implantId: String!
) {
  tasks(implantUuid: $implantId) {
    uuid
    type
    payload
    stdout
    stderr
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
    "concreteType": "Task",
    "kind": "LinkedField",
    "name": "tasks",
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
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "tasksGetTasksQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "tasksGetTasksQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "885a082811de5ad22bd3528e1e7c8c86",
    "id": null,
    "metadata": {},
    "name": "tasksGetTasksQuery",
    "operationKind": "query",
    "text": "query tasksGetTasksQuery(\n  $implantId: String!\n) {\n  tasks(implantUuid: $implantId) {\n    uuid\n    type\n    payload\n    stdout\n    stderr\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '5b2c57082f6d468e915d9ae958028292';

module.exports = node;
