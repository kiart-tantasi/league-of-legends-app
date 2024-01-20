import './App.css'
import { FormEvent, useContext, useRef, useState } from 'react'
import MatchContext from './contexts/MatchContext'

function App() {
  const { matches } = useContext(MatchContext)
  return <>{matches.length === 0 ? <SearchPage /> : <MatchesPage />}</>
}

export default App

// TODO: move to separate file
function SearchPage() {
  const { setMatches } = useContext(MatchContext)
  const [isLoading, setIsLoading] = useState(false)
  const nameRef = useRef<HTMLInputElement>(null)
  const tagRef = useRef<HTMLInputElement>(null)

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    // validate inputs
    const name = nameRef.current?.value
    const tag = tagRef.current?.value
    ;[name, tag].forEach((val) => {
      if (!val || typeof val !== 'string') {
        alert('กรอกข้อมูลไม่ครบถ้วนหรือไม่ถูกต้อง')
        return
      }
    })
    // fetch matches
    try {
      const url = process.env.REACT_APP_API_DOMAIN as string
      const response = await fetch(
        `${url}/api/v1/matches?gameName=${name}&tagLine=${tag}`,
      )
      if (response.status === 200) {
        const json = await response.json()
        setMatches(json.matches)
        console.log(json)
      } else {
        alert('ไม่พบ กรุณาลองใหม่')
      }
    } catch (e) {
      console.error(e)
    } finally {
      setIsLoading(false)
    }
  }

  if (isLoading) {
    return <div>loading...</div>
  }

  return (
    <div
      className="flex justify-center items-center w-full h-screen"
      data-testid="root-app"
    >
      <form
        className="flex flex-col w-screen max-w-[450px] py-4 items-center bg-blue-100 sm:rounded-xl"
        onSubmit={onSubmit}
      >
        <input
          type="text"
          id="name"
          placeholder="ชื่อ"
          className="w-[250px] mb-4"
          ref={nameRef}
        />
        <input
          type="text"
          id="tag"
          placeholder="tag"
          className="w-[250px] mb-4"
          ref={tagRef}
        />
        <button className="bg-white rounded p-1 text-[14px]" type="submit">
          ดู match history
        </button>
      </form>
    </div>
  )
}

// TODO: move to separate file
function MatchesPage() {
  const { matches } = useContext(MatchContext)
  return (
    <div className="flex flex-col justify-center  pt-4">
      {/* TEMP */}
      {matches.map((match) => {
        const { championName, kills, deaths, assists } = match
        const text = `${championName}: ${kills}/${deaths}/${assists}`
        return (
          <div className="mb-2" key={text.replaceAll(' ', '_')}>
            {text}
          </div>
        )
      })}
    </div>
  )
}
