import * as Comlink from "comlink";
import Worker from "worker-loader!./access.worker.js";

const worker = Comlink.wrap(new Worker());

export async function generateAccess(...args) {
	return worker.generateAccess(...args);
}

export default generateAccess;
