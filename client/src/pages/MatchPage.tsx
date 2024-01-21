import { useContext } from 'react'
import MatchContext from '../contexts/MatchContext'

export default function MatchPage() {
  const { matches, setMatches } = useContext(MatchContext)

  const onBack = () => {
    setMatches([])
  }

  return (
    <div className="flex flex-col justify-center pt-2">
      <button className="p-2 w-fit mb-4 border" type="button" onClick={onBack}>
        กลับ
      </button>
      {matches.map((match, index) => {
        const { championName, kills, deaths, assists } = match
        const text = `${championName}: ${kills}/${deaths}/${assists}`
        const backgroundColor = !!match.win ? 'bg-blue-100' : 'bg-red-100'
        return (
          <div
            className={`mb-2 bg-blue-100 max-w-[1280px] ${backgroundColor} p-2`}
            key={text.replaceAll(' ', '_').concat(index.toString())}
          >
            {text}
          </div>
        )
      })}
    </div>
  )
}
