import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { join } from 'path';

const YAML_CONFIG_FILENAME = 'config.yaml';
const filePath = join(__dirname, '../config/', YAML_CONFIG_FILENAME);

// 因为ConfigModule有一个load方法，它需要的是一个函数
export default () => {
  return yaml.load(readFileSync(filePath, 'utf8'));
};
