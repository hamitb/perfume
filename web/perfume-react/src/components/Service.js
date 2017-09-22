import { PerfumeServiceService } from '../perfume.pb';

const Service = new PerfumeServiceService('http://localhost:8081');

export default Service;