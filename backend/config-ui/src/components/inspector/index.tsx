/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import { Drawer } from 'antd';

import * as S from './styled';

interface Props {
  open: boolean;
  data: any;
  title?: string;
  onClose?: () => void;
}

export const Inspector = ({ open, data, title, onClose }: Props) => {
  return (
    <Drawer bodyStyle={{ padding: 0 }} open={open} title={title} onClose={onClose}>
      <S.Wrapper>
        <div className="title">
          <h3>JSON CONFIGURATION</h3>
          <span>application/json</span>
        </div>
        <p className="description">This is the configuration format used in a blueprint's advanced mode.</p>
        <div className="content">
          <code>
            <pre>{JSON.stringify(data, null, '  ')}</pre>
          </code>
        </div>
      </S.Wrapper>
    </Drawer>
  );
};
