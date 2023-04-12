export const CircleGenreInput = ({ circleGenre, setCircleGenre }) => {
  const handleChange = (e) => {
    setCircleGenre(e.target.value)
  }
  
  return (
    <div className="flex flex-col">
      <label className="text-lg font-bold">ジャンル</label>
      <div className="mt-2">
        <input
          type="text"
          value={circleGenre}
          onChange={handleChange}
          className="w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:border-blue-500"
        />
      </div>
    </div>
  )
}

