import { Environment, Network, RecordSource, Store } from 'relay-runtime';
import fetchGraphQL from './fetch-graphql';

async function fetchRelay(params, variables) {
  console.debug(`fetching query ${params.name} with ${JSON.stringify(variables)}`);
  return fetchGraphQL(params.text, variables);
}

export default new Environment({
  network: Network.create(fetchRelay),
  store: new Store(new RecordSource())
});
