import { createStore, combineReducers } from 'redux';
import taskReducer from './reducers/taskReducer';

const rootReducer = combineReducers({
  tasks: taskReducer,
});

const store = createStore(
  rootReducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

export default store;
