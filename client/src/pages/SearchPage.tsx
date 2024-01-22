import { FormEvent, useState } from 'react'
import { useNavigate } from 'react-router-dom'

export default function SearchPage() {
  const [gameName, setGameName] = useState('')
  const [tagLine, setTagLine] = useState('')
  const navigate = useNavigate()

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!validateSearchInputs({ gameName, tagLine })) {
      warnUser('กรอกข้อมูลไม่ครบ/ไม่ถูกต้อง')
      return
    }
    navigate(`/match?gameName=${gameName}&tagLine=${tagLine.replace('#', '')}`)
  }

  return (
    <div
      className="flex justify-center w-full h-fit mt-[200px]"
      data-testid="root-app"
    >
      <form
        className="flex flex-col w-screen max-w-[450px] py-4 items-center bg-blue-100"
        onSubmit={onSubmit}
      >
        <input
          value={gameName}
          onChange={(e) => setGameName(e.target.value)}
          type="text"
          name="gameName"
          placeholder="ชื่อในเกม"
          className="w-[250px] mb-4 placeholder:text-[0.75rem]"
        />
        <input
          value={tagLine}
          onChange={(e) => {
            setTagLine(e.target.value)
          }}
          type="text"
          name="tagLine"
          placeholder="#1234"
          className="w-[250px] mb-4 placeholder:text-[0.75rem]"
        />
        <button className="bg-white rounded p-1 text-[14px]" type="submit">
          ดู match history
        </button>
      </form>
    </div>
  )
}

export function handleTagLine({ tagLine }: { tagLine: string }) {
  return tagLine.replace('#', '')
}

export function validateSearchInputs({
  gameName,
  tagLine,
}: {
  gameName: string | null
  tagLine: string | null
}) {
  return [gameName, tagLine].every(
    (e) => typeof e === 'string' && e.length !== 0,
  )
}

export function warnUser(str: string) {
  // TODO: create shared warning modal
  alert(str)
}
