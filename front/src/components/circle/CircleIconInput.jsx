export const CircleIconInput = ({ circleIcon, setCircleIcon}) => {
  const handleChange = (e) => {
    console.log(e.target.files[0])
  }

  return (
    <div className="flex flex-col">
      <label htmlFor="image" className="text-lg font-bold mb-2">
        アイコン
      </label>
      <div className="mt-2">
        <label
          htmlFor="image"
          className="cursor-pointer bg-gray-100 rounded-md py-2 px-4 text-gray-700 font-medium hover:bg-gray-200"
        >
          画像を選択
        </label>
        <input
          type="file"
          id="image"
          onChange={handleChange}
          className="hidden"
        />
      </div>
    </div>
  )
}

