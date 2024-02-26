import React, { useEffect, useState } from "react";
import { PhotoIcon } from "@heroicons/react/24/outline";

import getPreview from "../../helpers/api/webApi/file/getPreview";
import ErrorToast from "../../helpers/toast/errorToast";

const File = ({ file }) => {
	const [preview, setPreview] = useState("");

	useEffect(() => {
		getPreview(file.id).then((d) => {
			setPreview(d.url);
		}).catch((error) => {
			ErrorToast(error);
		});
	}, []);

	return (
		<a
			className="border-[1px] border-gray-500 rounded-md flex p-4 my-2 shadow-md hover:underline"
			href={`/file/${file.id}`}
		>
			<div className="mr-8">
				{preview === ""
					? <PhotoIcon className="h-10 w-10" />
					: <img alt="preview" src={preview} className="w-16 h-auto" /> }
			</div>
			<div>
				<h2 className="font-semibold text-xl">{file.name}</h2>
				<div className="w-full flex !no-underline">
					<p className="text-gray-800 mr-4 !no-underline">Author: </p>
					<p className="font-medium !no-underline">{file.owner_name}</p>
				</div>
			</div>
		</a>
	);
};

export default File;
