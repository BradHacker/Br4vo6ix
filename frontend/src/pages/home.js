import React, { useCallback, useEffect, useState } from 'react';
import { graphql } from 'babel-plugin-relay/macro';
import { loadQuery, usePreloadedQuery, useQueryLoader } from 'react-relay/hooks';
import { commitMutation } from 'react-relay';

import LoadingScreen from '../components/loading';
import relayEnvironment from '../relay-environment';

import './home.css';

const { Suspense } = React;

const GetImplantsQuery = graphql`
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
`;

const NewTaskMutation = graphql`
  mutation homeNewTaskMutation($implantId: String!, $type: TaskType!, $payload: String!) {
    scheduleTask(input: { implantUuid: $implantId, type: $type, payload: $payload }) {
      uuid
    }
  }
`;

export const preloadedQuery = loadQuery(relayEnvironment, GetImplantsQuery, {});

const Home = (props) => {
  const [isExpanded, setIsExpanded] = useState([]);
  const data = usePreloadedQuery(GetImplantsQuery, props.preloadedQueryRef);

  const toggleExpansion = (implantId) => {
    let index = isExpanded.indexOf(implantId);
    if (index === -1) setIsExpanded([...isExpanded, implantId]);
    else {
      isExpanded.splice(index, 1);
      setIsExpanded([...isExpanded]);
    }
  };

  const checkExpanded = (implantId) => isExpanded.indexOf(implantId) >= 0;

  return (
    <table className="table">
      <thead>
        <tr>
          <th scope="col">machine_id</th>
          <th scope="col">last_seen_at</th>
          <th scope="col">controls</th>
        </tr>
      </thead>
      <tbody>
        {data.implants &&
          data.implants.map((implant, i) => (
            <React.Fragment key={`implant_${i}`}>
              <tr>
                <td scope="row">{implant.machine_id}</td>
                <td>{implant.last_seen_at}</td>
                <td>
                  <div className="btn-group" role="group" aria-label="Implant Controls">
                    <button type="button" className="btn btn-outline-secondary" onClick={() => toggleExpansion(implant.uuid)}>
                      <i className={`fa fa-chevron-${checkExpanded(implant.uuid) ? 'up' : 'down'}`}></i>
                    </button>
                    <button
                      type="button"
                      className="btn btn-outline-danger"
                      data-bs-toggle="modal"
                      data-bs-target="#exampleModal"
                      onClick={() => props.setSelectedImplant(implant.uuid)}
                    >
                      <i className="fa fa-terminal"></i>
                    </button>
                  </div>
                </td>
              </tr>
              {implant.tasks.length > 0 ? (
                implant.tasks.map((task, i) => (
                  <tr key={`implant_${implant.uuid}_task_${i}`} className={checkExpanded(implant.uuid) ? 'expanded' : 'collapsed'}>
                    <td colSpan="4" className="task-td">
                      <div className="task-row">
                        <div>{task.uuid}</div>
                        <div>{task.type}</div>
                        <div>{task.payload}</div>
                        <div>{task.stdout || task.stderr}</div>
                        <div>{task.created_at}</div>
                      </div>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td>No tasks :(</td>
                </tr>
              )}
            </React.Fragment>
          ))}
      </tbody>
    </table>
  );
};

const HomeRoot = (props) => {
  const [selectedImplant, setSelectedImplant] = useState('');
  const [implantCommand, setImplantCommand] = useState('');
  const [queryRes, setQueryRes] = useState(null);
  const [queryRef, loadQuery] = useQueryLoader(GetImplantsQuery, preloadedQuery);

  const refetch = useCallback(() => {
    loadQuery({}, { fetchPolicy: 'network-only' });
  }, []);

  useEffect(() => {
    // refresh the data every 30 secs
    const pid = setInterval(refetch, 30 * 1000);

    return () => clearInterval(pid);
  });

  const createNewTask = () => {
    return commitMutation(relayEnvironment, {
      mutation: NewTaskMutation,
      variables: {
        implantId: selectedImplant,
        type: 'CMD',
        payload: implantCommand
      },
      onCompleted: (response) => {
        console.log(response);
        setQueryRes({
          type: 'success',
          message: `Successfully scheduled task`
        });
        refetch();
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
    <div className="container-fluid p-4">
      <div className="modal fade" id="exampleModal" tabIndex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div className="modal-dialog">
          <div className="modal-content bg-dark">
            <div className="modal-header">
              <h5 className="modal-title" id="exampleModalLabel">
                Command for {selectedImplant}
              </h5>
              <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div className="modal-body">
              <label htmlFor="exampleFormControlTextarea1" className="form-label">
                Command
              </label>
              <textarea
                className="form-control"
                id="exampleFormControlTextarea1"
                rows="5"
                value={implantCommand}
                onChange={(event) => setImplantCommand(event.target.value)}
              ></textarea>
              {queryRes && (
                <div className={`alert alert-${queryRes.type}`} role="alert">
                  {queryRes.message}
                </div>
              )}
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-outline-secondary" data-bs-dismiss="modal" onClick={() => setImplantCommand('')}>
                Cancel
              </button>
              <button type="button" className="btn btn-outline-primary" onClick={createNewTask}>
                Submit
              </button>
            </div>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <h1>Home</h1>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <Suspense fallback={<LoadingScreen />}>
            <Home preloadedQueryRef={queryRef} setSelectedImplant={setSelectedImplant} />
          </Suspense>
        </div>
      </div>
    </div>
  );
};

export default HomeRoot;
