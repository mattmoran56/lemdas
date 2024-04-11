import React, { useEffect, useState } from "react";
import { ClockIcon, EyeIcon } from "@heroicons/react/24/outline";
import convertUnixTimestampToStringDate from "../../helpers/utils/datetime";

const Activity = ({
	activity, last,
}) => {
	const [details, setDetails] = useState({});
	useEffect(() => {
		setDetails(JSON.parse(activity.details));
	}, []);

	return (
		<div className="w-full flex">
			<div className="flex flex-col items-center">
				<div className="w-10 h-10 aspect-square bg-oxfordblue rounded-full flex items-center justify-center">
					{activity.type === "make_public"
						? <EyeIcon className="h-5 w-5 text-white" />
						: null}
				</div>
				{!last
					? <div className="h-12 w-1 mt-1 border-l-2 border-oxfordblue border-dashed" />
					: null}
			</div>
			<div className="flex-grow ml-4">
				{activity.type === "make_public"
					? (
						<p className="font-semibold mt-2">
							Published dataset: &nbsp;
							<a
								href={`/dataset/${details.dataset_id}`}
								className="text-oxfordblue-500 underline"
							>
								{details.dataset_name}
							</a>
						</p>
					) : null}
				<div className="mt-2 w-full flex items-center text-gray-500">
					<ClockIcon className="h-4 w-4 mr-2" />
					<p className="text-sm">on {convertUnixTimestampToStringDate(activity.created_at)}</p>
				</div>
			</div>
		</div>
	);
};

export default Activity;
