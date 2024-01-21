import { FormEvent, useContext, useRef, useState } from 'react'
import MatchContext from '../contexts/MatchContext'

export default function SearchPage() {
  const { setMatches } = useContext(MatchContext)
  const [isLoading, setIsLoading] = useState(false)
  const [name, setName] = useState('')
  const [tag, setTag] = useState('')
  const nameRef = useRef(null)
  const tagRef = useRef(null)

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault()
    // validate inputs
    const toValidates = [name, tag]
    toValidates.forEach((val) => {
      if (!val || typeof val !== 'string') {
        alert('กรอกข้อมูลไม่ครบถ้วน/ไม่ถูกต้อง')
        return
      }
    })
    // fetch matches
    try {
      setIsLoading(true)
      const response = await fetch(
        `/api/v1/matches?gameName=${name}&tagLine=${tag}`,
      )
      if (response.status === 200) {
        const json = await response.json()
        setMatches(json.matches)
        console.log(json)
      } else {
        alert('ไม่พบ')
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
          value={name}
          onChange={(e) => {
            setName(e.target.value)
          }}
          type="text"
          name="name"
          placeholder="ชื่อ"
          className="w-[250px] mb-4"
          ref={nameRef}
        />
        <input
          value={tag}
          onChange={(e) => {
            setTag(e.target.value)
          }}
          type="text"
          name="tag"
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
