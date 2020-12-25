import {Form, Modal, Select, notification} from 'antd';
import React, { useEffect, useState } from "react";
import api from "@/services/api";

interface ListFormProps {
    modalVisible: boolean;
    formTitle: string;
    initialValues: {};
    onSubmit: () => void;
    onCancel: () => void;
}

const formLayout = {
    labelCol: { span: 7 },
    wrapperCol: { span: 13 },
};

const ListForm: React.FC<ListFormProps> = (props) => {
    const { modalVisible, onCancel, onSubmit, initialValues, formTitle } = props;
    const [form] = Form.useForm();
    const [selectData, setSelectData] = useState([]); // 设置select

    useEffect(() => {
        if (form && !modalVisible) {
            form.resetFields();
        }
    }, [modalVisible]);

    useEffect(() => {
        if (initialValues) {
            form.setFieldsValue({
                ...initialValues,
            });
        }
    }, [initialValues]);

    // 选择select
    useEffect(() => {
        api.select().then((r)=>{
            if (r.code !== 0) {
                notification.error({
                    message: "加载失败",
                });
                return;
            }
            setSelectData(r.data)
        })
    }, []);

    const handleSubmit = () => {
        if (!form) return;
        form.submit();
    };

    const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel }
    return (
        <Modal
            destroyOnClose
            title={formTitle}
            visible={modalVisible}
            {...modalFooter}
        >
            <Form
                {...formLayout}
                form={form}
                onFinish={onSubmit}
                scrollToFirstError
            >
                <Form.Item
                    name="ids"
                    label="分类"
                >
                    <Select
                        style={{ width: '100%' }}
                        placeholder="分类"
                    >
                        {  selectData || []).map((item,index)=>{
                            return (<Option key={index} value={item.value}>{item.title}</Option>)
                        })}
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    );
};
export default ListForm;
