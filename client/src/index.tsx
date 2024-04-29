import ReactDOM from 'react-dom/client'
import './index.css'
import { MatchContextProvider } from './contexts/MatchContext'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import SearchPage from './pages/SearchPage'
import MatchPage from './pages/MatchPage'
// import reportWebVitals from './reportWebVitals';

const router = createBrowserRouter([
  {
    path: '*',
    element: <SearchPage />,
  },
  {
    path: '/match',
    element: <MatchPage />,
  },
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <MatchContextProvider>
    <RouterProvider router={router} />
  </MatchContextProvider>,
)

// TODO: enable web vitals

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();

// foobar
