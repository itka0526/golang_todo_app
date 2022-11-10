import { useState } from "react";

export default function UpDown({ mainRef }: { mainRef: React.RefObject<HTMLDivElement> }) {
    const [state, setState] = useState(0);
    const handleScroll = () => {
        mainRef.current?.children[state].scrollIntoView({ behavior: "smooth" });

        if (state === Number(mainRef.current?.children.length) - 1) {
            return setState(0);
        }
        return setState((prev) => prev + 1);
    };
    return (
        <div
            onClick={handleScroll}
            className="fixed bottom-7 right-7 border rounded-full shadow-md z-10 h-12 w-12 bg-white flex justify-center items-center "
        >
            <span className="text-4xl font-thin translate-y-1">^</span>
        </div>
    );
}
