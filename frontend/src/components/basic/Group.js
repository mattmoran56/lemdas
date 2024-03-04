import React from "react";
import { UsersIcon } from "@heroicons/react/24/solid";
import { MinusCircleIcon } from "@heroicons/react/24/outline";

const Group = ({
	name, access, onRemove, writeAccess, noAccess,
}) => {
	return (
		<div className="w-44 h-36 border-2 border-oxfordblue-200 my-2 mr-4 p-4 rounded-md flex flex-col items-center">
			<UsersIcon className="h-12 w-12 text-oxfordblue-400" />
			<p className="text-md font-semibold py-2">
				{name.length > 12 ? `${name.slice(0, 12)}...` : name}
			</p>
			<div className={`flex ${writeAccess ? "justify-between" : "justify-center"} w-full`}>
				{!noAccess ? <p className="text-gray-500 capitalize">{access}</p> : null}
				<button
					type="button"
					aria-label="remove user"
					onClick={onRemove}
					disabled={!writeAccess}
					className={writeAccess ? "" : "hidden"}
				>
					<MinusCircleIcon className="h-6 w-6 mr-3 text-gray-400 hover:text-red-500" />
				</button>
			</div>
		</div>
	);
};

export default Group;
