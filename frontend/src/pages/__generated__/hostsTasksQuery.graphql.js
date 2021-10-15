/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type TaskType = "CMD" | "NOOP" | "SCRIPT" | "%future added value";
export type hostsTasksQueryVariables = {|
  implantId: string
|};
export type hostsTasksQueryResponse = {|
  +tasks: $ReadOnlyArray<{|
    +uuid: string,
    +type: TaskType,
    +payload: string,
    +stdout: ?string,
    +stderr: ?string,
    +created_at: any,
  |}>
|};
export type hostsTasksQuery = {|
  variables: hostsTasksQueryVariables,
  response: hostsTasksQueryResponse,
|};
*/


/*
query hostsTasksQuery(
  $implantId: String!
) {
  tasks(implantUuid: $implantId) {
    uuid
    type
    payload
    stdout
    stderr
    created_at
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
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "hostsTasksQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "hostsTasksQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "439d5f76dd1ec55d3aa39a5af43aa0d9",
    "id": null,
    "metadata": {},
    "name": "hostsTasksQuery",
    "operationKind": "query",
    "text": "query hostsTasksQuery(\n  $implantId: String!\n) {\n  tasks(implantUuid: $implantId) {\n    uuid\n    type\n    payload\n    stdout\n    stderr\n    created_at\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = 'b7ed186518dfc8fdbb8b257efe7c76ab';

module.exports = node;
