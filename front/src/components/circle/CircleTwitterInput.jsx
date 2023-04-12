export const CircleTwitterInput = ({ circleTwitter, setCircleTwitter }) => {
  const handleChange = (e) => {
    setCircleTwitter(e.target.value)
  }
  
  return (
    <div className="flex flex-col">
      <label className="text-lg font-bold">twitter</label>
      <div className="mt-2">
        <input
          type="text"
          value={circleTwitter}
          onChange={handleChange}
          className="w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
        />
      </div>
    </div>
  )
}