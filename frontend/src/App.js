import React from 'react';
import { BrowserRouter as Router, Route, Switch, NavLink } from 'react-router-dom';
import logo from './logo.svg';
import './App.css';

import Home from './pages/home';
import { RelayEnvironmentProvider } from 'react-relay/hooks';
import relayEnvironment from './relay-environment';

/*
<header className="App-header">
<img src={logo} className="App-logo" alt="logo" />
    <p>
    Edit <code>src/App.js</code> and save to reload.
    </p>
    <a className="App-link" href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
    Learn React
    </a>
    </header>
    {data.implants && data.implants.map((implant, i) => <div key={`implant_${i}`}>{implant.uuid}</div>)}
    */

const App = (props) => {
  return (
    <RelayEnvironmentProvider environment={relayEnvironment}>
      <Router>
        <div className="App text-light">
          <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
            <div className="container-fluid">
              <a className="navbar-brand" href="#">
                Br4vo6ix
              </a>
              <button
                className="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent"
                aria-expanded="false"
                aria-label="Toggle navigation"
              >
                <span className="navbar-toggler-icon"></span>
              </button>
              <div className="collapse navbar-collapse" id="navbarSupportedContent">
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                  <li className="nav-item">
                    <NavLink className="nav-link" aria-current="page" to="/" exact>
                      Home
                    </NavLink>
                  </li>
                </ul>
              </div>
            </div>
          </nav>
          <Switch>
            <Route path="/" component={Home} />
          </Switch>
        </div>
      </Router>
    </RelayEnvironmentProvider>
  );
};

export default App;
