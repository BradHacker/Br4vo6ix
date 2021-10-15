import React, { useCallback, useEffect, useState } from 'react';
import { graphql } from 'babel-plugin-relay/macro';
import { loadQuery, usePreloadedQuery, useQueryLoader, useLazyLoadQuery } from 'react-relay/hooks';
import { commitMutation } from 'react-relay';
import { useParams } from 'react-router';

import LoadingScreen from '../components/loading';
import relayEnvironment from '../relay-environment';

import './hosts.css';

const { Suspense } = React;

const GetTasksQuery = graphql`
  query hostsTasksQuery($implantId: String!) {
    tasks(implantUuid: $implantId) {
      uuid
      type
      payload
      stdout
      stderr
      created_at
    }
  }
`;

const GetImplantQuery = graphql`
  query hostsImplantQuery($implantId: String!) {
    implant(implantUuid: $implantId) {
      uuid
      hostname
      ip
      machine_id
      last_seen_at
    }
  }
`;

const NewTaskMutation = graphql`
  mutation hostsNewTaskMutation($implantId: String!, $type: TaskType!, $payload: String!) {
    scheduleTask(input: { implantUuid: $implantId, type: $type, payload: $payload }) {
      uuid
    }
  }
`;

// export const preloadedQuery = loadQuery(relayEnvironment, GetTasksQuery, {});

const Hosts = (props) => {
  const { queryOptions, implantId, refresh } = props;
  const tasksData = useLazyLoadQuery(
    GetTasksQuery,
    {
      implantId
    },
    queryOptions
  );
  const [implantCommand, setImplantCommand] = useState('');
  const [queryRes, setQueryRes] = useState(null);
  // const data = usePreloadedQuery(GetTasksQuery, props.preloadedQueryRef);
  // const implant = usePreloadedQuery(GetImplantQuery, props.preloadedImplantQuery);
  // console.log(props);

  const createNewTask = () => {
    return commitMutation(relayEnvironment, {
      mutation: NewTaskMutation,
      variables: {
        implantId,
        type: 'CMD',
        payload: implantCommand
      },
      onCompleted: (response) => {
        // console.log(response);
        setQueryRes({
          type: 'success',
          message: `Successfully scheduled task`
        });
        refresh();
      },
      onError: (error) => {
        console.error(error);
        setQueryRes({
          type: 'danger',
          message: `Error scheduling task`
        });
      }
    });
  };

  return (
    <React.Fragment>
      <div className="row">
        <div className="col">
          <h1>
            Implant - {props.implantData.implant.hostname} {props.implantData.implant.ip}
          </h1>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <textarea
            className="form-control implant-screen"
            value={tasksData.tasks.map((task) => `>>>> ${task.payload}\n\n${task.stdout || task.stderr}\n\n`).join('\n')}
            readOnly
          ></textarea>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <label htmlFor="exampleFormControlTextarea1" className="form-label">
            Command
          </label>
          <textarea
            className="form-control"
            id="exampleFormControlTextarea1"
            rows="1"
            value={implantCommand}
            onChange={(event) => setImplantCommand(event.target.value)}
          ></textarea>
          {queryRes && (
            <div className={`alert alert-${queryRes.type}`} role="alert">
              {queryRes.message}
            </div>
          )}
          <button type="button" className="btn btn-block w-100 mt-2 btn-outline-danger" onClick={createNewTask}>
            Send Command
          </button>
        </div>
      </div>
    </React.Fragment>
  );
};

const HostsRoot = (props) => {
  let { implantId } = useParams();
  const [queryRes, setQueryRes] = useState(null);
  // const [queryRef, loadQuery] = useQueryLoader(GetTasksQuery, preloadedQuery);
  const implantData = useLazyLoadQuery(GetImplantQuery, { implantId });
  // const tasksData = useLazyLoadQuery(GetTasksQuery, { implantId });
  const [refreshedQueryOptions, setRefreshedQueryOptions] = useState(null);
  // const [implantQueryRef, loadImplantQuery] = useQueryLoader(GetImplantQuery, preloadedImplantQuery);

  // loadImplantQuery({
  //   implantId
  // });
  console.log(implantId);

  const refresh = useCallback(
    () => {
      setRefreshedQueryOptions((prev) => ({
        fetchKey: (prev?.fetchKey ?? 0) + 1,
        fetchPolicy: 'network-only'
      }));
    },
    [
      /* ... */
    ]
  );

  useEffect(() => {
    // refresh the data every 30 secs
    const pid = setInterval(refresh, 30 * 1000);

    return () => clearInterval(pid);
  });

  return (
    <div className="container-fluid p-4">
      <Suspense fallback={<LoadingScreen />}>
        <Hosts implantId={implantId} implantData={implantData} queryOptions={refreshedQueryOptions ?? {}} refresh={refresh} />
      </Suspense>
    </div>
  );
};

export default HostsRoot;
