import React from 'react';

import 'fomantic-ui-css/semantic.css';
import {
  Container,
  Grid,
  Segment,
  Header,
  Image,
  Menu,
  Input,
  Dropdown,
  Item
} from 'semantic-ui-react'
import './App.css';
import { Link, Route, Switch } from "react-router-dom";


const filterOptions = [
  {
    key: "author", value: "author", "text": "Author",
  },
  {
    key: "title", "value": "title", "text": "Title"
  }
]

const style = {
  h1: {
    marginTop: '3em',
  },
  h2: {
    margin: '4em 0em 2em',
  },
  h3: {
    marginTop: '2em',
    padding: '2em 0em',
  },
  last: {
    marginBottom: '300px',
  },
  borderless: {
    border: 'none',
    boxShadow: 'none'
  },
  // NOTE: I hate CSS.
  inputWidth: {
    width: '72%'
  }
}

const trendingBooks = [
  {
    title: "Eragon",
    cover: "https://upload.wikimedia.org/wikipedia/en/c/ce/Eragon_book_cover.png",
    author: "Christopher Paolini",
  },
  {
    title: "Eldest",
    cover: "https://upload.wikimedia.org/wikipedia/en/e/e0/Eldest_book_cover.png",
    author: "Christopher Paolini",
  },
  {
    title: "Brisingr",
    cover: "https://upload.wikimedia.org/wikipedia/en/7/70/Brisingr_book_cover.png",
    author: "Christopher Paolini",
  },
  {
    title: "Inheritance ",
    cover: "https://upload.wikimedia.org/wikipedia/en/2/2b/Inheritance2011.JPG",
    author: "Christopher Paolini",
  }
]

const FixedNavigationMenu = () => (
  <div>
    <Menu fixed='top' inverted>
      <Container>
        <Menu.Item header position="left">
          <Image size='mini' src='/logo.png' style={{ marginRight: '1.5em' }} />
          <Link to="/">Athenaeum</Link>
        </Menu.Item>

        <Menu position="right" inverted>
          <Menu.Item position='right' ><Link to="/book">Book</Link></Menu.Item>
          <Menu.Item position='right' ><Link to="/collection">Collections</Link></Menu.Item>
        </Menu>
      </Container>
    </Menu>
  </div >
);

// TODO: Make all items below into Home Component
const FilterSearchBar = () => (
  <Container>
    <Menu secondary stackable>
      <Menu.Item header>Filter By:</Menu.Item>
      <Menu.Item>
        <Dropdown
          placeholder=''
          search
          selection
          options={filterOptions}
        />
      </Menu.Item>
      <Menu.Item style={style.inputWidth}>
        <Input
          action={{ type: 'submit', content: 'Go' }}
          placeholder='Search'
        />
      </Menu.Item>

    </Menu>
  </Container >
)

const HomeHeaderComponent = () => (
  <header className="App-header">
    <Container>
      <Header className="yellow">
        Athenaeum
    </Header>
    </Container>
  </header>
)


export default function App() {
  return (
    <div>
      <FixedNavigationMenu />


      <Header
        as='h1'
        content='What We Currently Love'
        style={style.h1}
        textAlign='center'
      />

      <FilterSearchBar />

      <br />
      <br />

      <Grid container columns={3} stackable>

        {trendingBooks.map((book, index) => (
          <Grid.Column width={8}>
            <Item.Group link>
              <Item key={index}>
                <Item.Image size="medium" src={book.cover}></Item.Image>
                <Item.Content>
                  <Item.Header>{book.title}</Item.Header>
                  <h4>{book.author}</h4>
                </Item.Content>
              </Item>
            </Item.Group>
          </Grid.Column>
        )
        )}

      </Grid>


      <Switch>
        {/* <Route exact path="/" component={HomeHeaderComponent} /> */}
        {/* <Route path="/collection" component={} /> */}
      </Switch>
    </div>
  );
}

