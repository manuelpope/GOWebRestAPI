
import './App.css';
import {Switch,Route, BrowserRouter} from 'react-router-dom';
import ListDividers from './components/ListDividers';
import HomePage from './components/HomePage';
function App() {
  return (
    <div className="App">
      <h1>Hoola cuadrando los routers</h1>

      <BrowserRouter>
<Switch>
<Route exact path='/' component={HomePage}></Route>
<Route exact path='/Crops' component={ListDividers}></Route>
</Switch>


</BrowserRouter>

    </div>
  );
}

export default App;
