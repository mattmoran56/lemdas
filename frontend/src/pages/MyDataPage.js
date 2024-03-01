import React, { useEffect, useState } from "react";
import { StarIcon } from "@heroicons/react/24/solid";
import { ToastContainer } from "react-toastify";

import SearchBar from "../components/SearchBar/SearchBar";
import Dataset from "../components/basic/Dataset";
import getStaredDatasets from "../helpers/api/webApi/dataset/getStaredDatasets";
import ErrorToast from "../helpers/toast/errorToast";
import getDatasets from "../helpers/api/webApi/dataset/getDatasets";
import getSharedDatasets from "../helpers/api/webApi/dataset/getSharedDatasets";

const MyDataPage = () => {
	const [staredDatasets, setStaredDatasets] = useState([]);
	const [sharedDatasets, setSharedDatasets] = useState([]);
	const [myDatasets, setMyDatasets] = useState([]);

	useEffect(() => {
		getStaredDatasets().then((data) => {
			setStaredDatasets(data);
		}).catch((error) => {
			ErrorToast(error);
		});

		getSharedDatasets().then((data) => {
			setSharedDatasets(data);
		}).catch((error) => {
			ErrorToast(error);
		});

		getDatasets("created_at").then((data) => {
			setMyDatasets(data);
		}).catch((error) => {
			ErrorToast(error);
		});
	}, []);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<ToastContainer />
			<div className="w-full max-w-6xl m-auto p-8">
				<div className="mb-8">
					<h2 className="text-3xl font-semibold flex justify-between">
						Stared Datasets
						<StarIcon className="h-8 w-8 text-yellow-500 mr-2" />
					</h2>
					<div className="w-full h-[2px] bg-oxfordblue mb-4" />
					<div className="flex flex-wrap">
						{staredDatasets.length === 0
							? (
								<p className="ml-2">
									Your stared datasets will appear here!
								</p>
							)
							: null}
						{staredDatasets.map((dataset) => {
							return (
								<Dataset key={dataset.id} dataset={dataset} stared />
							);
						})}
					</div>
				</div>
				<div className="mb-8">
					<h2 className="text-3xl font-semibold flex justify-between">
						Shared with me
					</h2>
					<div className="w-full h-[2px] bg-oxfordblue mb-4" />
					<div className="flex flex-wrap">
						{sharedDatasets.length === 0
							? (
								<p className="ml-2">
									You&apos;ve not been shared any datasets!
								</p>
							)
							: null}
						{sharedDatasets.map((dataset) => {
							return (
								<Dataset key={dataset.id} dataset={dataset} />
							);
						})}
					</div>
				</div>
				<div className="mb-8">
					<h2 className="text-3xl font-semibold flex justify-between">
						My datasets
					</h2>
					<div className="w-full h-[2px] bg-oxfordblue mb-4" />
					<div className="flex flex-wrap">
						{myDatasets.length === 0
							? (
								<p className="ml-2">
									You&apos;ve got no datasets. Upload some scans to get started!
								</p>
							)
							: null}
						{myDatasets.map((dataset) => {
							return (
								<Dataset key={dataset.id} dataset={dataset} />
							);
						})}
					</div>
				</div>
			</div>

		</div>
	);
};

export default MyDataPage;
