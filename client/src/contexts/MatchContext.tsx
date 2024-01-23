import { ReactNode, createContext, useState } from 'react'
import { IMatch } from '../models/match'

interface IMatchContext {
  matches: IMatch[]
  setMatches: React.Dispatch<React.SetStateAction<IMatch[]>>
}

const MatchContext = createContext<IMatchContext>({
  matches: [],
  setMatches: () => {},
})

function MatchContextProvider({ children }: { children: ReactNode }) {
  const [matches, setMatches] = useState<IMatch[]>([])
  return (
    <MatchContext.Provider value={{ matches, setMatches }}>
      {children}
    </MatchContext.Provider>
  )
}

export default MatchContext

export { MatchContextProvider }
