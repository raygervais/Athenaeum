import React from 'react';

import 'fomantic-ui-css/semantic.css';
import {
  Container,
  Grid,
  Header,
  Image,
  Menu,
  Input,
  Dropdown,
  Item
} from 'semantic-ui-react'
import './App.css';
import { Link } from "react-router-dom";


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
    marginBottom: '0'
  },
  borderless: {
    border: 'none',
    boxShadow: 'none'
  },
  inputWidth: {
    minWidth: '72%'
  }
}

// NOTE: For Layout Mockup, all data to be replaced with API call to backend.
const trendingBooks = [
  {
    title: "Eragon",
    cover: "https://upload.wikimedia.org/wikipedia/en/c/ce/Eragon_book_cover.png",
    author: "Christopher Paolini",
    description: "Eragon is the first book in The Inheritance Cycle by American fantasy writer Christopher Paolini. Paolini, born in 1983, wrote the novel while still in his teens. After writing the first draft for a year, Paolini spent a second year rewriting and fleshing out the story and characters."
  },
  {
    title: "Eldest",
    cover: "https://upload.wikimedia.org/wikipedia/en/e/e0/Eldest_book_cover.png",
    author: "Christopher Paolini",
    description: "Eldest is the second novel in the Inheritance Cycle by Christopher Paolini and the sequel to Eragon. Eldest was first published in hardcover on August 23, 2005, and was released in paperback in September 2006. Eldest has been released in an audiobook format, and as an ebook. Wikipedia"
  },
  {
    title: "Brisingr",
    cover: "https://upload.wikimedia.org/wikipedia/en/7/70/Brisingr_book_cover.png",
    author: "Christopher Paolini",
    description: "Brisingr is the third novel in the Inheritance Cycle by Christopher Paolini. It was released on September 20, 2008. Originally, Paolini intended to conclude the then Inheritance Trilogy in three books, but during writing the third book he decided that the series was too complex to conclude in one book, because the single book would be close to 1,500 pages long."
  },
  {
    title: "Inheritance ",
    cover: "https://upload.wikimedia.org/wikipedia/en/2/2b/Inheritance2011.JPG",
    author: "Christopher Paolini",
    description: "The Inheritance Cycle was originally intended to be a trilogy, but Paolini has stated that during writing, the length of Brisingr grew, and the book was split into two parts to be published separately. Because of this, many plot elements originally intended for Brisingr are in Inheritance."
  }
]

const FixedNavigationMenu = () => (
  <Container>
    <Menu fixed='top' size="massive" borderless>
      <Menu.Item header position="left">
        <Image size='mini' src='/logo.png' style={{ marginRight: '1.5em' }} />
        <Link to="/">Athenaeum</Link>
      </Menu.Item>

      <Menu position="right" size="massive" secondary>
        <Menu.Item><Link to="/about">About</Link></Menu.Item>
        <Dropdown item text="Books">
          <Dropdown.Menu>
            <Dropdown.Item><Link to="/book/recommended">Recommend</Link></Dropdown.Item>
            <Dropdown.Item><Link to="/book/collection">Collection</Link></Dropdown.Item>
            <Dropdown.Item><Link to="/book/search">Search</Link></Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
        <Dropdown item text="Account">
          <Dropdown.Menu>
            <Dropdown.Item><Link to="/login">Login</Link></Dropdown.Item>
            <Dropdown.Item><Link to="/signup">Signup</Link></Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </Menu>
    </Menu>
  </Container>
);


// TODO: Make all items below into Home Component
export const FilterSearchBar = () => (
  <Container>
    <Menu secondary stackable size="massive">
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
          action={{ type: 'submit', content: 'Search' }}
          placeholder='Brisingr'
        />
      </Menu.Item>

    </Menu>
  </Container >
)

export default function App() {
  return (
    <div>
      <FixedNavigationMenu />


      <Header
        as='h1'
        content='What We Currently Love: The Inheritance Cycle'
        style={style.h1}
        textAlign='center'
      />

      <FilterSearchBar />

      <br />
      <br />

      <Grid container columns={2} stackable>

        {trendingBooks.map((book, index) => (
          <Grid.Column width={8} key={index}>
            <Item.Group link>
              <Item key={index}>
                <Item.Image size="small" src={book.cover}></Item.Image>
                <Item.Content>
                  <Item.Header>{book.title}</Item.Header>
                  <h4>{book.author}</h4>
                  <p>{book.description}</p>
                </Item.Content>
              </Item>
            </Item.Group>
          </Grid.Column>
        )
        )}

      </Grid>
    </div>

  );
}

