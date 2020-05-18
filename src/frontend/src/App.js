import React from 'react';
import './App.css';

// Example of how to import semantic-ui into react project.
import './semantic/dist/semantic.css'
import { Button } from 'semantic-ui-react';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        {/* This button is to show the implementation of semantic-UI for now.*/}
        <Button>Setting up Semantic-UI</Button>
        <p>Hello World!</p>
      </header>
    </div>
  );
}

export default App;
