/**
 * Copyright 2022 The Vitess Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

export const env: ()=> NodeJS.ProcessEnv = () => { 
    return {...window.env, ...process.env };
};

// process.env variables are always strings, hence this tiny helper function
// to transmute it into a boolean. It is a function, rather than a constant,
// to support dynamic updates to process.env in tests.
export const isReadOnlyMode = (): boolean => env().REACT_APP_READONLY_MODE === 'true';
