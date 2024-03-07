import React, { useState } from "react";
import { ArrowDownTrayIcon } from "@heroicons/react/24/outline";

import downloadFile from "../../helpers/api/download/downloadFile";
import ErrorToast from "../../helpers/toast/errorToast";
import Button from "../basic/Button";
import Loader from "../basic/Loader";

const DownloadFileButton = ({ file }) => {
	const [clicked, setClicked] = useState(false);
	const handleDownload = () => {
		setClicked(true);
		downloadFile(file.id).then((blob) => {
			const url = window.URL.createObjectURL(blob);
			const link = document.createElement("a");
			link.href = url;
			link.setAttribute("download", file.name);
			document.body.appendChild(link);
			link.click();
			link.remove();
			setClicked(false);
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	return (
		<Button
			className="mt-4 ml-4 !text-indianred bg-offwhite"
			onClick={handleDownload}
		>
			{clicked
				? <Loader className="h-3 w-3" outerClassName="m-0 mr-2 !p-1" />
				: <ArrowDownTrayIcon className="h-6 w-6 mr-2" />}
			Download File
		</Button>
	);
};

export default DownloadFileButton;
