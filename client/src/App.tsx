import './App.css'
import { useContext } from 'react'
import MatchContext from './contexts/MatchContext'
import SearchPage from './pages/SearchPage'
import MatchPage from './pages/MatchPage'

function App() {
  const { matches } = useContext(MatchContext)
  return <>{matches.length === 0 ? <SearchPage /> : <MatchPage />}</>
}

export default App
