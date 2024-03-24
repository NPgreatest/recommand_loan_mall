import numpy as np
import sys
from joblib import load
scaler =             load('python-script/scaler.joblib')
label_encoders = {
    "Gender":        load('python-script/label_encoder_Gender.joblib'),
    "Married":       load('python-script/label_encoder_Married.joblib'),
    "Dependents":    load('python-script/label_encoder_Dependents.joblib'),
    "Education":     load('python-script/label_encoder_Education.joblib'),
    "Self_Employed": load('python-script/label_encoder_Self_Employed.joblib'),
    "Property_Area": load('python-script/label_encoder_Property_Area.joblib'),
}
model = load('python-script/random_forest_model.joblib')


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

    # 进行预测
    prediction = model.predict(features_standard)
    return prediction


if __name__ == "__main__":
    if len(sys.argv) > 1:
        features = sys.argv[1:]
    else:
        features = ["Male", "Yes", "2", "Graduate", "Yes", 1320, 0, 100, 360, 1, "Urban"]

    prediction = predict_loan_status(features)
    if prediction == 1:
        print(1)
    else:
        print(0)