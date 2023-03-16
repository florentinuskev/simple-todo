import React from 'react'

const Fab = ({ onClick }: any) => {
    return (
        <button onClick={() => onClick}
            className="fixed z-90 bottom-10 right-8 bg-blue-300 w-20 h-20 rounded-full drop-shadow-lg flex justify-center items-center text-white text-4xl hover:bg-blue-500 hover:drop-shadow-2xl">
            +
        </button>
    )
}

export default Fab