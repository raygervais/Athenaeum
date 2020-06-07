import React from 'react';

import 'fomantic-ui-css/semantic.css';
import { Container, Header } from 'semantic-ui-react';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Container>
          <Header className="yellow">
            Athenaeum
          </Header>
        </Container>

        <p>Hello World!</p>
      </header>
    </div>
  );
}

export default App;
