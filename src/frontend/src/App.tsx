import { useEffect, useState } from "react";

interface Task {
    id: string;
    status: string;
    taskName: string;
}

function App() {
    const [createValue, setCreateValue] = useState("");
    const [list, setList] = useState([]);

    useEffect(() => {
        readHandler();
    }, []);

    const readHandler = async () => {
        const pending_response = await fetch("/api/tasks");
        const response = await pending_response.json();

        setList(response);
    };

    const createHandler = async (): Promise<void> => {
        const options: RequestInit = {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                message: createValue,
            }),
        };
        const pending_response = await fetch("/api/tasks", options);
        const response = await pending_response.json();

        setList(response);
    };

    const Update = async ({ id, status }: { id: string; status: boolean }) => {
        const options: RequestInit = {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                id: id,
                status: status ? "complete" : "not complete",
            }),
        };
        const pending_response = await fetch("/api/tasks", options);
        const response = await pending_response.json();

        setList(response);
    };

    const Delete = async ({ id }: { id: string }) => {
        const options: RequestInit = {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                id: id,
            }),
        };
        const pending_response = await fetch("/api/tasks", options);
        const response = (await pending_response.json()) || [];

        setList(response);
    };

    return (
        <main className="h-screen w-screen bg-slate-100 grid grid-cols-4">
            <section className="border-r bg-white ">
                <h3 className="flex justify-center items-center w-full h-12 mb-1">Create</h3>
                <form
                    className="w-full h-full flex gap-4 flex-col"
                    onSubmit={(e) => {
                        e.preventDefault();
                        createHandler();
                    }}
                >
                    <div className="h-10 w-full px-2">
                        <input className="w-full border  h-full" onChange={(e) => setCreateValue(e.target.value)} value={createValue}></input>
                    </div>
                    <div className="h-10 w-full px-2">
                        <button
                            onClick={(e) => {
                                e.preventDefault();
                                createHandler();
                            }}
                            className="w-full h-full shadow-md border rounded-md"
                        >
                            Add
                        </button>
                    </div>
                </form>
            </section>
            <section className="border-r bg-white  ">
                <h3 className="flex justify-center items-center w-full h-12 mb-1">Read</h3>
                <List list={list} enabled={true} update={() => null} type={"none"} />
            </section>
            <section className="border-r bg-white  ">
                <h3 className="flex justify-center items-center w-full h-12 ">Update</h3>
                <List list={list} update={Update} type={"Update"} />
            </section>
            <section className="border-r bg-white  ">
                <h3 className="flex justify-center items-center w-full h-12 mb-1">Delete</h3>
                <List list={list} update={Delete} type={"Delete"} />
            </section>
        </main>
    );
}

const List = ({ list, enabled = false, update, type }: { list: Task[]; enabled?: boolean; update: Function; type: string }) => {
    return (
        <div className="h-full w-full px-2">
            <ul className="h-full w-full">
                {list.map((task: Task) => (
                    <ListItem key={task.id} task={task} enabled={enabled} update={update} type={type} />
                ))}
            </ul>
        </div>
    );
};

const ListItem = ({ task, enabled, update, type }: { task: Task; enabled: boolean; update: Function; type: string }) => {
    const { id, status, taskName } = task;
    const [state, setState] = useState(false);

    useEffect(() => {
        setState(!status.includes("not"));
    }, [status]);

    return (
        <li className={`w-full h-10 grid ${enabled ? "grid-cols-[90%,10%]" : "grid-cols-[60%,10%,30%]"} items-center`}>
            <span>{taskName}</span>
            <input type={"checkbox"} disabled={type === "Delete" || enabled} checked={state} onChange={() => setState((prev) => !prev)} />
            {!enabled && <button onClick={() => update({ id, status: state })}> {type} </button>}
        </li>
    );
};

export default App;
