import { ReactNode } from 'react'

const TEMP_MAX_WIDTH = 'max-w-[600px]'

export default function Layout({
  children,
  background,
}: {
  children: ReactNode
  background?: 'BLUE' | 'WHITE'
}) {
  const backgroundClass = background === 'BLUE' ? 'bg-blue-100' : 'bg-white'
  return (
    <div className={`w-full h-screen ${backgroundClass}`}>
      <div className={`w-full ${TEMP_MAX_WIDTH} h-screen mx-auto`}>
        {children}
      </div>
    </div>
  )
}
