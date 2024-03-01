import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import { TrashIcon, StarIcon, PlusIcon } from "@heroicons/react/24/outline";
import { StarIcon as StarSolidIcon } from "@heroicons/react/24/solid";

import getDataset from "../helpers/api/webApi/dataset/getDataset";
import getDatasetFiles from "../helpers/api/webApi/file/getFilesByDataset";
import SearchBar from "../components/SearchBar/SearchBar";
import Loader from "../components/basic/Loader";
import File from "../components/basic/File";
import getDatasetAttributes from "../helpers/api/webApi/datasetAttributes/getDatasetAttributes";
import Attributes from "../components/DatasetPage/Attribute";
import Upload from "../components/IndexPage/Upload";
import ErrorToast from "../helpers/toast/errorToast";
import Button from "../components/basic/Button";
import deleteDataset from "../helpers/api/webApi/dataset/deleteDataset";
import updateDataset from "../helpers/api/webApi/dataset/updateDataset";
import getStaredDataset from "../helpers/api/webApi/dataset/getStaredDataset";
import updateStaredDataset from "../helpers/api/webApi/dataset/updateStaredDataset";
import getCollaborators from "../helpers/api/webApi/datasetCollaborators/getCollaborators";
import CollaboratorsModal from "../components/DatasetPage/CollaboratorsModal";
import Sharing from "../components/DatasetPage/Sharing";

const DatasetPage = () => {
	const [dataset, setDataset] = useState({});
	const [files, setFiles] = useState([]);
	const [attributes, setAttributes] = useState([]);
	const [stared, setStared] = useState(false);
	const [collaborators, setCollaborators] = useState([]);
	const [writeAccess, setWriteAccess] = useState(false);
	const [modalOpen, setModalOpen] = useState(false);

	const [refreshAttribute, setRefreshAttribute] = useState(false);

	const { datasetId } = useParams();
	const navigate = useNavigate();

	const handleUpdateStared = () => {
		updateStaredDataset(datasetId).then(() => {
			getStaredDataset(datasetId).then((data) => {
				setStared(data.stared);
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	const handleUpdatePublic = (e) => {
		updateDataset(datasetId, dataset.dataset_name, e.target.checked, dataset.owner_id).then(() => {
			getDataset(datasetId)
				.then((data) => {
					setDataset(data);
				}).catch((error) => {
					ErrorToast(error);
				});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	const handleDelete = () => {
		deleteDataset(datasetId).then(() => {
			navigate("");
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		if (refreshAttribute) {
			getDatasetAttributes(datasetId).then((data) => {
				setAttributes(data);
			}).catch((error) => {
				ErrorToast(error);
			});
			setRefreshAttribute(false);
		}
	}, [refreshAttribute]);

	useEffect(() => {
		getDataset(datasetId)
			.then((data) => {
				setDataset(data);
				setWriteAccess(false);
			}).catch((error) => {
				ErrorToast(error);
			});
		getDatasetFiles(datasetId)
			.then((data) => {
				setFiles(data);
			}).catch((error) => {
				ErrorToast(error);
			});
		getDatasetAttributes(datasetId)
			.then((data) => {
				setAttributes(data);
			}).catch((error) => {
				ErrorToast(error);
			});
		getStaredDataset(datasetId)
			.then((data) => {
				setStared(data.stared);
			}).catch((error) => {
				ErrorToast(error);
			});
		getCollaborators(datasetId).then((data) => {
			setCollaborators(data.collaborators);
		}).catch((error) => {
			ErrorToast(error);
		});
		setRefreshAttribute(false);
	}, []);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<ToastContainer />
			{writeAccess
				? (
					<CollaboratorsModal
						datasetId={datasetId}
						isOpen={modalOpen}
						setIsOpen={setModalOpen}
						onClose={() => {
							getCollaborators(datasetId).then((data) => {
								setCollaborators(data.collaborators);
							}).catch((error) => {
								ErrorToast(error);
							});
						}}
					/>
				) : null }
			{ dataset
				? (
					<div className="flex justify-center items-center w-full">
						<div className="w-full p-8 max-w-7xl flex">
							<div className="w-1/2 p-2">
								<div className="w-full flex justify-between">
									<h1 className="text-3xl font-bold break-words">{dataset.dataset_name}</h1>
									{stared
										? (
											<button
												type="button"
												onClick={handleUpdateStared}
												aria-label="unstar dataset"
											>
												<StarSolidIcon className="h-6 w-6 text-yellow-500" />
											</button>
										)
										: (
											<button
												type="button"
												onClick={handleUpdateStared}
												aria-label="star dataset"
											>
												<StarIcon className="h-6 w-6 text-gray-400" />
											</button>
										)}
								</div>
								<div className="h-[2px] w-full bg-oxfordblue mb-4" />
								<div className="w-full flex">
									<p className="text-gray-800 mr-4">Author: </p>
									<p className="font-medium">{dataset.owner_name}</p>
								</div>
								<div className="w-full flex">
									<p className="text-gray-800 mr-4">Collaborators: </p>
									<p className="font-medium">
										{
											collaborators.map((collaborator, i) => {
												return (
													i === collaborators.length - 1
														? `${collaborators.length === 1 ? "" : "and"} 
														${collaborator.user.first_name} 
														${collaborator.user.last_name}`
														: `${collaborator.user.first_name} 
														${collaborator.user.last_name}, `
												);
											})
										}
										<button
											type="button"
											onClick={() => { setModalOpen(true); }}
											className={`ml-2 ${writeAccess ? "" : "hidden"}`}
											aria-label="show collaborators"
											disabled={!writeAccess}
										>
											<PlusIcon className="h-4 w-4 text-oxfordblue-400 mt-1" />
										</button>
									</p>
								</div>
								<div className="w-full flex items-center">
									<p className="text-gray-800 mr-4">Public dataset: </p>
									<p className="font-medium">
										<input
											type="checkbox"
											checked={dataset.is_public}
											onChange={handleUpdatePublic}
											disabled={!writeAccess}
										/>
									</p>
								</div>

								<Button
									className={`mt-4 ${writeAccess ? "" : "hidden"}`}
									onClick={handleDelete}
								>
									<TrashIcon className="h-6 w-6 mr-2" />
									Delete Dataset
								</Button>

								<h2 className="mt-6 text-2xl font-semibold">Dataset Attributes</h2>
								<div className="w-full max-h-64">
									<Attributes
										attributes={attributes}
										datasetId={datasetId}
										setNeedRefresh={setRefreshAttribute}
										writeAcces={writeAccess}
									/>
								</div>

								<div className="w-full h-[2px] bg-oxfordblue mt-8" />
								<h2 className="mt-6 text-2xl font-semibold">Sharing</h2>
								<Sharing datasetId={datasetId} writeAccess={writeAccess} />
							</div>
							<div className="w-1/2 p-2">
								<div className="w-full flex flex-wrap">
									{files.map((file) => {
										return (
											<File key={file.id} file={file} />
										);
									})}
								</div>
								<div className="w-full h-64 mt-8">
									{writeAccess
										? (
											<Upload
												datasetId={datasetId}
												onFinish={() => {
													getDatasetFiles(datasetId)
														.then((data) => {
															setFiles(data);
														}).catch((error) => {
															ErrorToast(error);
														});
												}}
											/>
										) : null}
								</div>
							</div>
						</div>
					</div>
				)
				: <Loader />}
		</div>
	);
};

export default DatasetPage;
