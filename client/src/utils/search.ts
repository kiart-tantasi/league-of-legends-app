export function handleTagLine(tagLine: string) {
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
