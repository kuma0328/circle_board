
export const CircleMemberInput = ({ circleMember, setCircleMember }) => {
  const handleChange = (value) => {
    setCircleMember(value)
  }
  
  return (
    <div className="flex flex-col">
      <label className="text-lg font-bold mb-2">人数</label>
      <div className="relative h-8 rounded-full bg-gray-200">
        <div
          className="absolute left-0 top-0 bottom-0 bg-blue-500 rounded-full"
          style={{ width: `${(circleMember / 500) * 100}%` }}
        />
        <input
          type="range"
          min="1"
          max="500"
          value={circleMember}
          onChange={(e) => handleChange(e.target.value)}
          className="absolute left-0 top-0 w-full h-full opacity-0 cursor-pointer"
        />
      </div>
      <p className="text-gray-500">{circleMember}</p>
    </div>
  )
}