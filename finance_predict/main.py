import numpy as np
import sys
from joblib import load

# 加载预处理模型和随机森林模型
scaler = load('scaler.joblib')
label_encoders = {
    "Gender": load('label_encoder_Gender.joblib'),
    "Married": load('label_encoder_Married.joblib'),
    "Dependents": load('label_encoder_Dependents.joblib'),
    "Education": load('label_encoder_Education.joblib'),
    "Self_Employed": load('label_encoder_Self_Employed.joblib'),
    "Property_Area": load('label_encoder_Property_Area.joblib'),
}
model = load('random_forest_model.joblib')


def predict_loan_status(features):
    features_transformed = list(features)
    features_transformed[0] = label_encoders["Gender"].transform([features[0]])[0]
    features_transformed[1] = label_encoders["Married"].transform([features[1]])[0]
    features_transformed[2] = label_encoders["Dependents"].transform([features[2]])[0]
    features_transformed[3] = label_encoders["Education"].transform([features[3]])[0]
    features_transformed[4] = label_encoders["Self_Employed"].transform([features[4]])[0]
    features_transformed[10] = label_encoders["Property_Area"].transform([features[10]])[0]

    # 将特征数组转换为numpy数组，并进行标准化
    features_array = np.array(features_transformed, dtype=object).reshape(1, -1)
    features_standard = scaler.transform(features_array.astype(float))
    prediction = model.predict(features_standard)
    return prediction


if __name__ == "__main__":
    if len(sys.argv) > 1:
        features = sys.argv[1:]
    else:
        features = ["Male", "Yes", "0", "Graduate", "No", 1000, 0, 10000, 360, 1, "Urban"]

    prediction = predict_loan_status(features)
    # 假设标签0和1分别代表贷款不批准和批准
    print("贷款状态预测:", "批准" if prediction == 1 else "不批准")


# python predict_loan_status.py Male Yes 0 Graduate No 3000 0 130 360 1 Urban