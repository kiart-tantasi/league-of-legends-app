import { FormEvent, useContext, useEffect, useState } from 'react'
import MatchContext from '../contexts/MatchContext'
import { Link, useSearchParams } from 'react-router-dom'
import {
  handleTagLine,
  validateSearchInputs,
  warnUser,
} from './../utils/search'

export default function MatchPage() {
  const [searchParams, setSearchParams] = useSearchParams()
  const paramGameName = searchParams.get('gameName')
  const paramTagLine = searchParams.get('tagLine')
  const [gameName, setGameName] = useState('')
  const [tagLine, setTagLine] = useState('')
  const { matches, setMatches } = useContext(MatchContext)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    if (
      !validateSearchInputs({ gameName: paramGameName, tagLine: paramTagLine })
    ) {
      warnUser('ไม่พบชื่อ/tagline')
      setIsLoading(false)
      return
    }

    ;(async () => {
      try {
        setIsLoading(true)
        const response = await fetch(
          `/api/v1/matches?gameName=${paramGameName}&tagLine=${handleTagLine({
            tagLine: paramTagLine ?? '',
          })}`,
        )
        if (response.status === 200) {
          const json = await response.json()
          setMatches(json.matches)
        } else {
          setMatches([])
          warnUser('ไม่พบ กรุณาลองใหม่')
        }
      } catch (e) {
        console.error(e)
      } finally {
        setIsLoading(false)
      }
    })()
  }, [paramGameName, paramTagLine, setMatches])

  const onSubmit = (e: FormEvent) => {
    e.preventDefault()
    if (
      [gameName, tagLine].some((e) => typeof e !== 'string' || e.length === 0)
    ) {
      warnUser('กรอกข้อมูลไม่ครบ/ไม่ถูกต้อง')
      return
    }
    setSearchParams({
      gameName,
      tagLine,
    })
  }

  if (isLoading) {
    return (
      <div className="text-center pt-[200px]">กำลังโด้ข้อมูลจาก RIOT...</div>
    )
  }
  return (
    <div className="flex flex-col justify-center pt-2 w-full max-w-[600px]">
      <div className="flex justify-between mb-4">
        <Link className="p-2 w-fit h-fit mb-4 border" type="button" to="/">
          กลับ
        </Link>
        <form className="flex flex-col w-[150px]" onSubmit={onSubmit}>
          <input
            value={gameName}
            onChange={(e) => setGameName(e.target.value)}
            placeholder="ชื่อในเกม"
            type="text"
            className="text-right mb-2"
          />
          <input
            value={tagLine}
            onChange={(e) => setTagLine(e.target.value)}
            placeholder="#1234"
            type="text"
            className="text-right mb-2"
          />
          <button type="submit" className="border">
            ดู match history
          </button>
        </form>
      </div>
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
