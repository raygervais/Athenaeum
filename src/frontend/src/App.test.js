import React from 'react'
import { BrowserRouter as Router, Route, MemoryRouter } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { render, fireEvent } from '@testing-library/react'
import App, { FilterSearchBar } from './App'

test('Application Renders to Homescreen', () => {
  const { container, getByText } = render(<App />, { wrapper: MemoryRouter })

  expect(container).not.toBeNull()
})

test('Application Renders to Filter', () => {
  const { container, getByText } = render(<App />, { wrapper: MemoryRouter })

  expect(container.innerHTML).toContain('Search')
  expect(container.innerHTML).toContain('Filter')
})
