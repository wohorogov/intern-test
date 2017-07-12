import { createStore, applyMiddleware } from 'redux';
import * as reducers from '../reducers/search-reducer';
import data from '../data/words'

let preloadedState = {
  words: data
};


const store = createStore(reducer, preloadedState);


export { preloadedState };
export default store;
