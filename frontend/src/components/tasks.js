import React from 'react';
import PropTypes from 'prop-types';
import { RelayEnvironmentProvider } from 'react-relay/hooks';
import { loadQuery } from 'react-relay/hooks';
import { usePreloadedQuery } from 'react-relay/hooks';
import relayEnvironment from '../relay-environment';
import { graphql } from 'babel-plugin-relay/macro';

const { Suspense } = React;

const GetTasksQuery = graphql`
  query tasksGetTasksQuery($implantId: String!) {
    tasks(implantUuid: $implantId) {
      uuid
      type
      payload
      stdout
      stderr
    }
  }
`;

export const preloadedQuery = loadQuery(relayEnvironment, GetTasksQuery, {});

const Tasks = (props) => {
  const data = usePreloadedQuery(GetTasksQuery, props.preloadedQuery, {
    variables: {
      implantId: props.implantUuid
    }
  });
  return (
    <React.Fragment>
      {data.tasks &&
        data.tasks.map((task, i) => {
          <tr key={`implant_${props.implantUuid}_task_${i}`}>
            <td>{task.uuid}</td>
            <td>{task.type}</td>
            <td>{task.payload}</td>
            <td>{task.stdout}</td>
            <td>{task.stderr}</td>
          </tr>;
        })}
    </React.Fragment>
  );
};

Tasks.defaultProps = {
  preloadedQuery: PropTypes.func,
  implantUuid: PropTypes.string
};

const TasksRoot = (props) => {
  return (
    <RelayEnvironmentProvider environment={relayEnvironment}>
      <Suspense
        fallback={
          <tr className="text-center">
            <i className="fa fa-spinner fa-spinner"></i>
          </tr>
        }
      >
        <Tasks preloadedQuery={preloadedQuery} implantUuid={props.implantUuid} />
      </Suspense>
    </RelayEnvironmentProvider>
  );
};

TasksRoot.defaultProps = {
  implantUuid: PropTypes.string
};

export default TasksRoot;
