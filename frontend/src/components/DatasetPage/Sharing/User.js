import React from "react";
import { UserIcon } from "@heroicons/react/24/solid";
import { MinusCircleIcon } from "@heroicons/react/24/outline";

const User = ({
	name, avatar, access, onRemove,
}) => {
	return (
		<div className="w-44 h-36 border-2 border-oxfordblue-200 my-2 mr-4 p-4 rounded-md flex flex-col items-center">
			{avatar
				? (
					<div className="aspect-square overflow-hidden h-24 rounded-full">
						<img src={avatar} alt="avatar" className="w-full h-auto mb-2 " />
					</div>
				)
				: <UserIcon className="h-12 w-12 text-oxfordblue-400" />}
			<p className="text-md font-semibold py-2">
				{name.length > 12 ? `${name.slice(0, 12)}...` : name}
			</p>
			<div className="flex justify-between w-full">
				<p className="text-gray-500 capitalize">{access}</p>
				<button
					type="button"
					aria-label="remove user"
					onClick={onRemove}
				>
					<MinusCircleIcon className="h-6 w-6 mr-3 text-gray-400 hover:text-red-500" />
				</button>
			</div>
		</div>
	);
};

export default User;
