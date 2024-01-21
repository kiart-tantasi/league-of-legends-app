import { ReactNode, createContext, useState } from 'react'

interface IMatch {
  championName: string
  kills: number
  deaths: number
  assists: number
  win: boolean
  gameMode: string
  gameCreation: number
}

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
