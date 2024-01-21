import { useContext } from 'react'
import MatchContext from '../contexts/MatchContext'

export default function MatchPage() {
  const { matches, setMatches } = useContext(MatchContext)

  const onBack = () => {
    setMatches([])
  }

  return (
    <div className="flex flex-col justify-center pt-2 w-full max-w-[600px]">
      <button className="p-2 w-fit mb-4 border" type="button" onClick={onBack}>
        กลับ
      </button>
      {matches.map((match, index) => {
        const { championName, kills, deaths, assists, gameMode } = match
        const backgroundColor = !!match.win ? 'bg-blue-100' : 'bg-red-100'
        return (
          <div
            className={`mb-2 p-2 ${backgroundColor}`}
            key={index.toString().concat('-match-detail')}
          >
            <div className="flex flex-row justify-between">
              <p>{championName}</p>
              <p>{`${kills}/${deaths}/${assists}`}</p>
            </div>
            <p className="text-[0.7rem]">{gameMode}</p>
            <p className="text-[0.7rem]">
              {new Date(match.gameCreation).toLocaleDateString('pt-PT')}
            </p>
          </div>
        )
      })}
    </div>
  )
}
